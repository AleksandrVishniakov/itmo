package pokemon;

import move.BatonPass;
import move.DoubleTeam;
import move.FakeOut;
import move.MuddyWater;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.Type;

final public class Scizor extends Pokemon {
  public Scizor(String name, int level) {
    super(name, level);
    this.setType(Type.BUG, Type.STEEL);
    this.setStats(70, 130, 100, 55, 80, 65);
    this.setMove(
      new MuddyWater(),
      new FakeOut(),
      new DoubleTeam(),
      new BatonPass()
    );
  }
}
