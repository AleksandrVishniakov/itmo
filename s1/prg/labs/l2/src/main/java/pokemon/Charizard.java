package pokemon;

import move.Growth;
import ru.ifmo.se.pokemon.Type;

final public class Charizard extends Charmeleon {
  public Charizard(String name, int level) {
    super(name, level);
    this.setType(Type.FIRE, Type.FLYING);
    this.setStats(78, 84, 78, 109, 85, 100);
    this.addMove(new Growth());
  }
}
