package pokemon;

import move.Flatter;

public class Charmeleon extends Charmander {
  public Charmeleon(String name, int level) {
    super(name, level);
    this.setStats(58, 64, 58, 80, 65, 80);
    this.addMove(new Flatter());
  }
}
